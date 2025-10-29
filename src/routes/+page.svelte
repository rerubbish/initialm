<script>
	export let data;
	import { onMount } from 'svelte';
	import Go from '../lib/wsam/wasm_exec';
	import InputFiled from '../lib/component/inputFiled.svelte';

	console.log('加载的JSON数据:', data);
	/**
	 * @type {Go | undefined}
	 */
	let go = undefined;
	/**
	 * @type {WebAssembly.Module | undefined}
	 */
	let mod = undefined;
	/**
	 * @type {WebAssembly.WebAssemblyInstantiatedSource | WebAssembly.Instance | undefined}
	 */
	let inst = undefined;

	async function runWsam() {
		if (!inst) {
			console.error('inst未初始化，无法运行');
			return;
		}
		if (!mod) {
			console.error('mod未初始化，无法运行');
			return;
		}
		if (!go) {
			console.error('Go未初始化，无法运行');
			return;
		}
		await go.run(inst);
		inst = await WebAssembly.instantiate(mod, go.importObject); // reset instance
	}

	onMount(() => {
		go = new Go();
		if (!WebAssembly.instantiateStreaming) {
			// polyfill
			WebAssembly.instantiateStreaming = async (resp, importObject) => {
				const source = await (await resp).arrayBuffer();
				return await WebAssembly.instantiate(source, importObject);
			};
		}

		WebAssembly.instantiateStreaming(fetch('wasm/app.wasm'), go.importObject)
			.then((result) => {
				mod = result.module;
				inst = result.instance;
				console.log('wsam加载成功');
				runWsam();
				// 调用add函数（传递数字参数）
				const sum = window.add(10, 20);
				console.log('10 + 20 =', sum); // 输出：30
			})
			.catch((err) => {
				console.error(err);
			});
	});
</script>

<div>
	<div>
		<h1 class="text-3xl font-bold">application initializr</h1>
		<form>
			{#each data.rulesJsonData[0].data.components as component}
				{#if component.type === 'text'}
					<InputFiled data={component} />
				{/if}
			{/each}
			<button class="btn" type="submit">Submit form</button>
		</form>
	</div>
</div>
