<script>
	export let data;
</script>

<legend class="fieldset-legend">{data.label}</legend>
<div>
	{#each data.options as option}
		<label
			><input
				name={data.key}
				class="checkbox validator"
				type="checkbox"
				value={option.value}
				required={option.required || false}
				placeholder={data.tips}
			/>{option.lable}</label
		>
		<span class="validator-hint">
			{@html option.required && '必填 <br />'}
		</span>
	{/each}
</div>
