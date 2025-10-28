<script>
	export let data;
</script>

<fieldset class="fieldset">
	<legend class="fieldset-legend">{data.label}</legend>
	<input
		type="text"
		class="input validator"
		required={data.validate.required}
		placeholder={data.tips}
		pattern={data.validate.ValidatorExpr.rule}
	/>
	<span class="validator-hint">
		{@html data.validate.required && '必填 <br />'}
		{data.validate.ValidatorExpr.message}
	</span>
</fieldset>
