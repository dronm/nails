/** Copyright (c) 2023
	Andrey Mikhalevich, Katren ltd.
 *
 * Номер счета включает двадцать знаков
 * Первые 5 знаков
 * 	-40817 - Текущие и карточные счета физических лиц
 * 	-42301 - Депозиты до востребования 
 * Далее идет три знака с обозначением валюты счета: 
 * 	-810 - российский рубль (RUR),
 * 	-840 - американский доллар (USD),
 * 	-978 - евро (EUR) и т.д.
 * Далее идет один знак с контрольной суммой
 * Далее идет четыре цифры с кодом отделения (филиала) банка
 * И наконец, последние 7 цифр - собственно номер счета
 */
function BankAccEdit(id,options){
	options = options || {};
	options.labelCaption = (options.labelCaption!=undefined)? options.labelCaption : "Расчетный счет:";
	options.title = options.title || "Расчетный счет";	
	options.placeholder = options.placeholder || "Расчетный счет";
	options.maxLength = 20;
	options.regExpression = "^[0-9]{20}$";
	/*
	options.formatterGetRawValue = true;
	options.formatterOptions = {
		"blocks": [5, 3, 1, 4, 7],
		"delimiter": " ",
		"numeral": true,
		"numeralThousandsGroupStyle":"none"
	};	
	*/
	
	BankAccEdit.superclass.constructor.call(this,id,options);	
}
extend(BankAccEdit, EditNum);//string

BankAccEdit.prototype.validate = function(){
	if(!EditINNSelfEmployed.superclass.validate.call(this)){
		return false;
	}
	if(this.m_bik && this.m_bik.length && this.m_bik.length == 9 && !fn_checkRS(this.getValue(), this.m_bik)){
		this.setNotValid(this.ER_REGEXP_INVALID);
		return false;
	}
	return true;
}

BankAccEdit.prototype.setBik = function(bik){
	this.m_bik = bik
}

//функция проверки правильности указания банковского счёта
function fn_bank_account(s){        
	var result = false;
	var sm = 0;

	//весовые коэффициенты
	var v = [7,1,3,7,1,3,7,1,3,7,1,3,7,1,3,7,1,3,7,1,3,7,1];
	for (var i = 0; i <= 22; i++){
		//вычисляем контрольную сумму
		sm = sm + ( Number(s.charAt(i)) * v[i] ) % 10;
	}

	//сравниваем остаток от деления контрольной суммы на 10 с нулём
console.log("fn_bank_account sm=" + sm)	
	if(sm % 10 == 0){
		result = true;
	}
	return result;         
}

function fn_checkRS(accNum, bik){
	//return fn_bank_account(bik.substr(6,3) + accNum);
	return fn_bank_account(bik.substring(6,9) + accNum);
}

