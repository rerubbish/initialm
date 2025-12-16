<script>
	export let data;
	import { onMount } from 'svelte';
	import Go from '$lib/wsam/wasm_exec';
	import InputField from '$lib/component/InputField.svelte';
	import FileInputField from '$lib/component/FileInputField.svelte';
	import SelectField from '$lib/component/SelectField.svelte';
	import CheckboxField from '$lib/component/CheckboxField.svelte';

	const { rulesData = [] } = data;
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

<div class="h-full w-full">
	<div class="mt-4 m-auto w-xs space-y-4">
		<h1 class="text-3xl font-bold">application initializr</h1>
		<div>
			<legend class="fieldset-legend">选择模板</legend>
			<select name="ruleName" onchange={ruleNameSelectChange} class="select mb-8">
				{#each rulesData as data, index}
					<option value={JSON.stringify(data.data)} selected={index === 0}>
						{data.data?.name || data.filename}
					</option>
				{/each}
			</select>
		</div>
		<form class="w-xs">
			<fieldset class="fieldset">
				{#each genformData.components as component}
					{#if component.type === 'text'}
						<InputField data={component} />
					{/if}
					{#if component.type === 'file'}
						<FileInputField data={component} />
					{/if}
					{#if component.type === 'select'}
						<SelectField data={component} />
					{/if}
					{#if component.type === 'checkbox'}
						<CheckboxField data={component} />
					{/if}
				{/each}
				<button class="btn btn-primary" type="submit">Submit form</button>
			</fieldset>
		</form>
	</div>
</div>
