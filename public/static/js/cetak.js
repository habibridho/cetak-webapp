$(function () {
    const inputTempaltePath = $('#inputTemplatePath');
    const inputTemplate = $('#inputTemplateFile');
    const uploadTemplateButton = $('#uploadTemplate');

    if (inputTempaltePath) {
        inputTempaltePath.click(function () {
            inputTemplate.trigger('click');
        });
    }

    inputTemplate.change(function () {
        inputTempaltePath.text(inputTemplate.val());
        if (inputTemplate.val() === "") {
            uploadTemplateButton.removeClass('hidden');
            inputTempaltePath.addClass('hidden');
        } else {
            uploadTemplateButton.addClass('hidden');
            inputTempaltePath.removeClass('hidden');
        }
    });

    if (uploadTemplateButton) {
        uploadTemplateButton.click(function () {
            inputTemplate.trigger('click');
        });
    }
});