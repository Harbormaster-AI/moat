import React, { Component } from 'react'
import UoMConversionService from '../services/UoMConversionService';

class UpdateUoMConversionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                factor: '',
                precision: '',
                fromUnit: '',
                toUnit: ''
        }
        this.updateUoMConversion = this.updateUoMConversion.bind(this);

        this.changefactorHandler = this.changefactorHandler.bind(this);
        this.changeprecisionHandler = this.changeprecisionHandler.bind(this);
        this.changeFromUnitHandler = this.changeFromUnitHandler.bind(this);
        this.changeToUnitHandler = this.changeToUnitHandler.bind(this);
    }

    componentDidMount(){
        UoMConversionService.getUoMConversionById(this.state.id).then( (res) =>{
            let uoMConversion = res.data;
            this.setState({
                factor: uoMConversion.factor,
                precision: uoMConversion.precision,
                fromUnit: uoMConversion.fromUnit,
                toUnit: uoMConversion.toUnit
            });
        });
    }

    updateUoMConversion = (e) => {
        e.preventDefault();
        let uoMConversion = {
            uoMConversionId: this.state.id,
            factor: this.state.factor,
            precision: this.state.precision,
            fromUnit: this.state.fromUnit,
            toUnit: this.state.toUnit
        };
        console.log('uoMConversion => ' + JSON.stringify(uoMConversion));
        console.log('id => ' + JSON.stringify(this.state.id));
        UoMConversionService.updateUoMConversion(uoMConversion).then( res => {
            this.props.history.push('/uoMConversions');
        });
    }

    changefactorHandler= (event) => {
        this.setState({factor: event.target.value});
    }
    changeprecisionHandler= (event) => {
        this.setState({precision: event.target.value});
    }
    changeFromUnitHandler= (event) => {
        this.setState({fromUnit: event.target.value});
    }
    changeToUnitHandler= (event) => {
        this.setState({toUnit: event.target.value});
    }

    cancel(){
        this.props.history.push('/uoMConversions');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update UoMConversion</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> factor: </label>
                                                <input placeholder="factor" name="factor" className="form-control" value={this.state.factor} onChange={this.changefactorHandler}/>

                                            <label> precision: </label>
                                                <input type="number" placeholder="precision" name="precision" className="form-control" value={this.state.precision} onChange={this.changeprecisionHandler}/>

                                            <label> FromUnit: </label>
                                                <select value={this.state.fromUnit} onChange={this.changeFromUnitHandler}>
                      <option name="FromUnit" className="form-control" >
                          Each
                      </option>
                      <option name="FromUnit" className="form-control" >
                          Case
                      </option>
                      <option name="FromUnit" className="form-control" >
                          Pallet
                      </option>
                      <option name="FromUnit" className="form-control" >
                          Dozen
                      </option>
                      <option name="FromUnit" className="form-control" >
                          Gram
                      </option>
                      <option name="FromUnit" className="form-control" >
                          Kilogram
                      </option>
                      <option name="FromUnit" className="form-control" >
                          Pound
                      </option>
                      <option name="FromUnit" className="form-control" >
                          Ounce
                      </option>
                      <option name="FromUnit" className="form-control" >
                          Milliliter
                      </option>
                      <option name="FromUnit" className="form-control" >
                          Liter
                      </option>
                      <option name="FromUnit" className="form-control" >
                          CubicMeter
                      </option>
                      <option name="FromUnit" className="form-control" >
                          Meter
                      </option>
                      <option name="FromUnit" className="form-control" >
                          Foot
                      </option>
                      <option name="FromUnit" className="form-control" >
                          SquareMeter
                      </option>
                    </select>

                                            <label> ToUnit: </label>
                                                <select value={this.state.toUnit} onChange={this.changeToUnitHandler}>
                      <option name="ToUnit" className="form-control" >
                          Each
                      </option>
                      <option name="ToUnit" className="form-control" >
                          Case
                      </option>
                      <option name="ToUnit" className="form-control" >
                          Pallet
                      </option>
                      <option name="ToUnit" className="form-control" >
                          Dozen
                      </option>
                      <option name="ToUnit" className="form-control" >
                          Gram
                      </option>
                      <option name="ToUnit" className="form-control" >
                          Kilogram
                      </option>
                      <option name="ToUnit" className="form-control" >
                          Pound
                      </option>
                      <option name="ToUnit" className="form-control" >
                          Ounce
                      </option>
                      <option name="ToUnit" className="form-control" >
                          Milliliter
                      </option>
                      <option name="ToUnit" className="form-control" >
                          Liter
                      </option>
                      <option name="ToUnit" className="form-control" >
                          CubicMeter
                      </option>
                      <option name="ToUnit" className="form-control" >
                          Meter
                      </option>
                      <option name="ToUnit" className="form-control" >
                          Foot
                      </option>
                      <option name="ToUnit" className="form-control" >
                          SquareMeter
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateUoMConversion}>Save</button>
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

export default UpdateUoMConversionComponent
