const circle = require('./helpers/circle');

exports.handler = function(context, event) {
    context.callback(""+circle.area(6));
};
