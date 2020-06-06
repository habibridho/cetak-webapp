$(function () {
    const uploadSelectors = [{
        inputPath: '#inputTemplatePath',
        input: '#inputTemplateFile',
        uploadButton: '#uploadTemplate',
    }, {
        inputPath: '#inputDataPath',
        input: '#inputDataFile',
        uploadButton: '#uploadData',
    }]

    uploadSelectors.forEach(function (selector) {
        const inputPath = $(selector.inputPath);
        const input = $(selector.input);
        const uploadButton = $(selector.uploadButton);

        if (inputPath) {
            inputPath.click(function () {
                input.trigger('click');
            });
        }

        input.change(function () {
            inputPath.text(input.val());
            if (input.val() === "") {
                uploadButton.removeClass('hidden');
                inputPath.addClass('hidden');
            } else {
                uploadButton.addClass('hidden');
                inputPath.removeClass('hidden');
            }
        });

        if (uploadButton) {
            uploadButton.click(function () {
                input.trigger('click');
            });
        }
    });
});