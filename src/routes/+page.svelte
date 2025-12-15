<script>
	export let data;
	const { rulesData = [] } = data;
	import { onMount } from 'svelte';
	import Go from '../lib/wsam/wasm_exec';
	import InputFiled from '../lib/component/inputFiled.svelte';
	$: genformData = rulesData[0].data;
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

	function ruleNameSelectChange(e) {
		const rule = JSON.parse(e.target.value);
		genformData = rule;
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
			})
			.catch((err) => {
				console.error(err);
			});
	});
</script>

<div>
	<div>
		<h1 class="text-3xl font-bold">application initializr</h1>
		<select name="ruleName" onchange={ruleNameSelectChange}>
			<option value="[]">无</option>
			{#each rulesData as data, index}
				<option value={JSON.stringify(data.data)} selected={index === 0}> {data.filename} </option>
			{/each}
		</select>
		<form>
			{#each genformData.components as component}
				{#if component.type === 'text'}
					<InputFiled data={component} />
				{/if}
			{/each}
			<button class="btn" type="submit">Submit form</button>
		</form>
	</div>
</div>
