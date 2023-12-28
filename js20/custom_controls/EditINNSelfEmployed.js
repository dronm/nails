/** Copyright (c) 2023
	Andrey Mikhalevich, Katren ltd.
 */
function EditINNSelfEmployed(id,options){
	options = options || {};
	options.labelCaption = (options.labelCaption!=undefined)? options.labelCaption : "ИНН:";
	options.title = options.title || "ИНН самозанятого";	
	options.placeholder = options.placeholder || "ИНН самозанятого";
	options.maxLength = 12;
	options.regExpression = "^[0-9]{12}$";
	
	EditINNSelfEmployed.superclass.constructor.call(this,id,options);	
}
extend(EditINNSelfEmployed, EditNum);

EditINNSelfEmployed.prototype.validate = function(){
	if(!EditINNSelfEmployed.superclass.validate.call(this)){
		return false;
	}
	if(!CheckINN_FL(this.getValue())){
		this.setNotValid(this.ER_REGEXP_INVALID);
		return false;
	}
	return true;
}

// Валидатор ИНН

function CheckINN_FL(inn){
	var summ1 =
		(
			(
				07 * parseInt(inn.charAt(00)) +
				02 * parseInt(inn.charAt(01)) +
				04 * parseInt(inn.charAt(02)) +
				10 * parseInt(inn.charAt(03)) +
				03 * parseInt(inn.charAt(04)) +
				05 * parseInt(inn.charAt(05)) +
				09 * parseInt(inn.charAt(06)) +
				04 * parseInt(inn.charAt(07)) +
				06 * parseInt(inn.charAt(08)) +
				08 * parseInt(inn.charAt(09))
			) % 11
		) % 10;

	var summ2 =
		(
			(
				03 * parseInt(inn.charAt(00)) +
				07 * parseInt(inn.charAt(01)) +
				02 * parseInt(inn.charAt(02)) +
				04 * parseInt(inn.charAt(03)) +
				10 * parseInt(inn.charAt(04)) +
				03 * parseInt(inn.charAt(05)) +
				05 * parseInt(inn.charAt(06)) +
				09 * parseInt(inn.charAt(07)) +
				04 * parseInt(inn.charAt(08)) +
				06 * parseInt(inn.charAt(09)) +
				08 * parseInt(inn.charAt(10))
			) % 11
		) % 10;

	return (parseInt(inn.charAt(10)) == summ1 && parseInt(inn.charAt(11)) == summ2);
}


