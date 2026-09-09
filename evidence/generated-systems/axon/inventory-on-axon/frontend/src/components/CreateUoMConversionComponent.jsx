import React, { Component } from 'react'
import UoMConversionService from '../services/UoMConversionService';

class CreateUoMConversionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                factor: '',
                precision: '',
                fromUnit: '',
                toUnit: ''
        }
        this.changefactorHandler = this.changefactorHandler.bind(this);
        this.changeprecisionHandler = this.changeprecisionHandler.bind(this);
        this.changeFromUnitHandler = this.changeFromUnitHandler.bind(this);
        this.changeToUnitHandler = this.changeToUnitHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateUoMConversion = (e) => {
        e.preventDefault();
        let uoMConversion = {
                uoMConversionId: this.state.id,
                factor: this.state.factor,
                precision: this.state.precision,
                fromUnit: this.state.fromUnit,
                toUnit: this.state.toUnit
            };
        console.log('uoMConversion => ' + JSON.stringify(uoMConversion));

        // step 5
        if(this.state.id === '_add'){
            uoMConversion.uoMConversionId=''
            UoMConversionService.createUoMConversion(uoMConversion).then(res =>{
                this.props.history.push('/uoMConversions');
            });
        }else{
            UoMConversionService.updateUoMConversion(uoMConversion).then( res => {
                this.props.history.push('/uoMConversions');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add UoMConversion</h3>
        }else{
            return <h3 className="text-center">Update UoMConversion</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> factor:&emsp; </label>
                                                <input placeholder="factor" name="factor" className="form-control" value={this.state.factor} onChange={this.changefactorHandler}/>

                                            <label> precision:&emsp; </label>
                                                <input type="number" placeholder="precision" name="precision" className="form-control" value={this.state.precision} onChange={this.changeprecisionHandler}/>

                                            <label> FromUnit:&emsp; </label>
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

                                            <label> ToUnit:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateUoMConversion}>Save</button>
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

export default CreateUoMConversionComponent
