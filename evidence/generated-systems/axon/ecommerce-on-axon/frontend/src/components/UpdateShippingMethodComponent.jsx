import React, { Component } from 'react'
import ShippingMethodService from '../services/ShippingMethodService';

class UpdateShippingMethodComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                flatRate: '',
                estimatedDays: '',
                asActive: '',
                methodType: ''
        }
        this.updateShippingMethod = this.updateShippingMethod.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeflatRateHandler = this.changeflatRateHandler.bind(this);
        this.changeestimatedDaysHandler = this.changeestimatedDaysHandler.bind(this);
        this.changeasActiveHandler = this.changeasActiveHandler.bind(this);
        this.changeMethodTypeHandler = this.changeMethodTypeHandler.bind(this);
    }

    componentDidMount(){
        ShippingMethodService.getShippingMethodById(this.state.id).then( (res) =>{
            let shippingMethod = res.data;
            this.setState({
                name: shippingMethod.name,
                flatRate: shippingMethod.flatRate,
                estimatedDays: shippingMethod.estimatedDays,
                asActive: shippingMethod.asActive,
                methodType: shippingMethod.methodType
            });
        });
    }

    updateShippingMethod = (e) => {
        e.preventDefault();
        let shippingMethod = {
            shippingMethodId: this.state.id,
            name: this.state.name,
            flatRate: this.state.flatRate,
            estimatedDays: this.state.estimatedDays,
            asActive: this.state.asActive,
            methodType: this.state.methodType
        };
        console.log('shippingMethod => ' + JSON.stringify(shippingMethod));
        console.log('id => ' + JSON.stringify(this.state.id));
        ShippingMethodService.updateShippingMethod(shippingMethod).then( res => {
            this.props.history.push('/shippingMethods');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeflatRateHandler= (event) => {
        this.setState({flatRate: event.target.value});
    }
    changeestimatedDaysHandler= (event) => {
        this.setState({estimatedDays: event.target.value});
    }
    changeasActiveHandler= (event) => {
        this.setState({asActive: event.target.value});
    }
    changeMethodTypeHandler= (event) => {
        this.setState({methodType: event.target.value});
    }

    cancel(){
        this.props.history.push('/shippingMethods');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ShippingMethod</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> flatRate: </label>
                                                <input placeholder="flatRate" name="flatRate" className="form-control" value={this.state.flatRate} onChange={this.changeflatRateHandler}/>

                                            <label> estimatedDays: </label>
                                                <input type="number" placeholder="estimatedDays" name="estimatedDays" className="form-control" value={this.state.estimatedDays} onChange={this.changeestimatedDaysHandler}/>

                                            <label> asActive: </label>
                                                <input type="checkbox" placeholder="asActive" name="asActive" className="form-control" value={this.state.asActive} onChange={this.changeasActiveHandler}/>


                                            <label> MethodType: </label>
                                                <select value={this.state.methodType} onChange={this.changeMethodTypeHandler}>
                      <option name="MethodType" className="form-control" >
                          Standard
                      </option>
                      <option name="MethodType" className="form-control" >
                          Expedited
                      </option>
                      <option name="MethodType" className="form-control" >
                          Overnight
                      </option>
                      <option name="MethodType" className="form-control" >
                          SameDay
                      </option>
                      <option name="MethodType" className="form-control" >
                          Pickup
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateShippingMethod}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>

                   </div>
            </div>
        )
    }
}

export default UpdateShippingMethodComponent
