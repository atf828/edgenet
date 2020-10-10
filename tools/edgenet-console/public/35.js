(window["webpackJsonp"] = window["webpackJsonp"] || []).push([[35],{

/***/ "./resources/js/data/order/OrderableContextNative.js":
/*!***********************************************************!*\
  !*** ./resources/js/data/order/OrderableContextNative.js ***!
  \***********************************************************/
/*! exports provided: OrderableContext, OrderableConsumer */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "OrderableContext", function() { return OrderableContext; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "OrderableConsumer", function() { return OrderableConsumer; });
/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! react */ "./node_modules/react/index.js");
/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(react__WEBPACK_IMPORTED_MODULE_0__);
function _typeof(obj) { "@babel/helpers - typeof"; if (typeof Symbol === "function" && typeof Symbol.iterator === "symbol") { _typeof = function _typeof(obj) { return typeof obj; }; } else { _typeof = function _typeof(obj) { return obj && typeof Symbol === "function" && obj.constructor === Symbol && obj !== Symbol.prototype ? "symbol" : typeof obj; }; } return _typeof(obj); }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } }

function _createClass(Constructor, protoProps, staticProps) { if (protoProps) _defineProperties(Constructor.prototype, protoProps); if (staticProps) _defineProperties(Constructor, staticProps); return Constructor; }

function _inherits(subClass, superClass) { if (typeof superClass !== "function" && superClass !== null) { throw new TypeError("Super expression must either be null or a function"); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, writable: true, configurable: true } }); if (superClass) _setPrototypeOf(subClass, superClass); }

function _setPrototypeOf(o, p) { _setPrototypeOf = Object.setPrototypeOf || function _setPrototypeOf(o, p) { o.__proto__ = p; return o; }; return _setPrototypeOf(o, p); }

function _createSuper(Derived) { return function () { var Super = _getPrototypeOf(Derived), result; if (_isNativeReflectConstruct()) { var NewTarget = _getPrototypeOf(this).constructor; result = Reflect.construct(Super, arguments, NewTarget); } else { result = Super.apply(this, arguments); } return _possibleConstructorReturn(this, result); }; }

function _possibleConstructorReturn(self, call) { if (call && (_typeof(call) === "object" || typeof call === "function")) { return call; } return _assertThisInitialized(self); }

function _assertThisInitialized(self) { if (self === void 0) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return self; }

function _isNativeReflectConstruct() { if (typeof Reflect === "undefined" || !Reflect.construct) return false; if (Reflect.construct.sham) return false; if (typeof Proxy === "function") return true; try { Date.prototype.toString.call(Reflect.construct(Date, [], function () {})); return true; } catch (e) { return false; } }

function _getPrototypeOf(o) { _getPrototypeOf = Object.setPrototypeOf ? Object.getPrototypeOf : function _getPrototypeOf(o) { return o.__proto__ || Object.getPrototypeOf(o); }; return _getPrototypeOf(o); }


var ReactContext = react__WEBPACK_IMPORTED_MODULE_0___default.a.createContext({});
var OrderableProvider = ReactContext.Provider;
var OrderableConsumer = ReactContext.Consumer;

var OrderableContext = /*#__PURE__*/function (_Component) {
  _inherits(OrderableContext, _Component);

  var _super = _createSuper(OrderableContext);

  function OrderableContext(props) {
    var _this;

    _classCallCheck(this, OrderableContext);

    _this = _super.call(this, props);
    _this.state = {
      mouseover: false,
      draggable: false,
      elementOver: false
    };
    _this.dragStart = _this.dragStart.bind(_assertThisInitialized(_this));
    _this.dragEnd = _this.dragEnd.bind(_assertThisInitialized(_this));
    _this.dragOver = _this.dragOver.bind(_assertThisInitialized(_this));
    return _this;
  }

  _createClass(OrderableContext, [{
    key: "dragStart",
    value: function dragStart(ev) {
      //console.log(ev);
      // this.dragged = Number(ev.currentTarget.dataset.id);
      ev.dataTransfer.effectAllowed = 'move'; // Firefox requires calling dataTransfer.setData
      // for the drag to properly work

      ev.dataTransfer.setData("text/html", null);
    }
  }, {
    key: "dragOver",
    value: function dragOver(ev, id) {
      ev.preventDefault();

      if (id !== this.state.elementOver) {
        // const items = this.state.list;
        var over = ev.currentTarget;
        console.log(id);
        this.setState({
          elementOver: id
        }); // const dragging = this.state.dragging;
        // const from = isFinite(dragging) ? dragging : this.dragged;
        // let to = Number(over.dataset.id);
        // items.splice(to, 0, items.splice(from,1)[0]);
        //console.log(over)
        //this.sort(items, to);
      }
    }
  }, {
    key: "dragEnd",
    value: function dragEnd(ev) {
      this.setState({
        elementOver: false
      }); //console.log(ev)
      // this.sort(this.state.list, undefined);
    }
  }, {
    key: "render",
    value: function render() {
      var children = this.props.children;
      var elementOver = this.state.elementOver;
      return /*#__PURE__*/react__WEBPACK_IMPORTED_MODULE_0___default.a.createElement(OrderableProvider, {
        value: {
          onDragStart: this.dragStart,
          onDragOver: this.dragOver,
          onDragEnd: this.dragEnd,
          elementOver: elementOver
        }
      }, children);
    }
  }]);

  return OrderableContext;
}(react__WEBPACK_IMPORTED_MODULE_0__["Component"]);



/***/ })

}]);