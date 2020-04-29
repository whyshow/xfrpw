var UserInfo=function(){
	
	 var userFormValidation = function() {
		var userForm = $('#userupdataForm');
		var error1 = $('.alert-danger', userForm);
	    var success1 = $('.alert-success', userForm);
		 userForm.validate({
                errorElement: 'span', //default input error message container
                errorClass: 'help-block help-block-error', // default input error message class
                focusInvalid: false, // do not focus the last invalid input
                ignore: "",  // validate all fields including form hidden input
                messages: {
                   
                },
                rules: {
                    teachername: {
                      
                        required: true
                    },
                    jobnum:{
                    	required: true,
                    	minlength:5
                    },
                    birth:{
                    	required: true
                    },
                    tel:{
                    	required: true
                    },
                    address:{
                    	required: true
                    },
                    graduate: {
                        required: true,
                        
                    },
                    inschool: {
                        required: true,
                       
                    },
                    direction: {
                        required: true,
                       
                    }
                    
                },

                invalidHandler: function (event, validator) { //display error alert on form submit              
                    success1.hide();
                    error1.show();
                    Metronic.scrollTo(error1, -200);
                },

                highlight: function (element) { // hightlight error inputs
                    $(element)
                        .closest('.form-group').addClass('has-error'); // set error class to the control group
                },

                unhighlight: function (element) { // revert the change done by hightlight
                    $(element)
                        .closest('.form-group').removeClass('has-error'); // set error class to the control group
                },

                success: function (label) {
                    label
                        .closest('.form-group').removeClass('has-error'); // set success class to the control group
                },

                submitHandler: function (form) {
                	
                     $.ajax({
		                 url :$("#url").val(),
		                 type :"POST",
		                 data : {
		                 	id:form.userid.value,
		                 	teachername:form.teachername.value,
		                 	Jobnum:form.jobnum.value,
		                 	sex:form.sex.value,
		                 	birth:form.birth.value,
		                 	tel:form.tel.value,
		                 	address:form.address.value,
		                 	graduate:form.graduate.value,
		                 	job:form.job.value,
		                 	working:form.working.value,
		                 	profession:form.profession.value,
		                 	check:form.check.value,
		                 	situation:form.situation.value,
		                 	abroad:form.abroad.value,
		                 	inschool:form.inschool.value,
		                 	direction:form.direction.value,
		                 	research:form.research.value,
		                 	jobtitle:form.jobtitle.value,
		                 	award:form.award.value,
		                 	paper:form.paper.value,
		                 	patent:form.patent.value
		                 },
	                	success : function(data) {
	                	console.log(data);
	                   
	                    	if(data==1){
								layer.alert('修改成功', {
									icon: 1,
									skin: 'layer-ext-moon', //该皮肤由layer.seaning.com友情扩展。关于皮肤的扩展规则，去这里查阅
									time:2000,
									end:function(){
										window.location.href=$("#objurl").val();
									}
									})
								
	                     		}else if(data==2){
	                     			layer.alert('修改失败，该工号已存在', {
									icon: 2,
									skin: 'layer-ext-moon', //该皮肤由layer.seaning.com友情扩展。关于皮肤的扩展规则，去这里查阅
									time:2000
									
									})
	                     		}else{
	                     			layer.alert('修改失败', {
									icon: 2,
									skin: 'layer-ext-moon', //该皮肤由layer.seaning.com友情扩展。关于皮肤的扩展规则，去这里查阅
									time:2000
									})
	                     	}
                			 }
           				  })
                }
	        });
	    }
	 
	 return{
	 	init: function() {
	 		userFormValidation();
        }
	 };
}();
