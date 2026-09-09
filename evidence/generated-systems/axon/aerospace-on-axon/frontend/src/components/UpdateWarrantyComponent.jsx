import React, { Component } from 'react'
import WarrantyService from '../services/WarrantyService';

class UpdateWarrantyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                coverageMonths: '',
                warrantyType: ''
        }
        this.updateWarranty = this.updateWarranty.bind(this);

        this.changecoverageMonthsHandler = this.changecoverageMonthsHandler.bind(this);
        this.changeWarrantyTypeHandler = this.changeWarrantyTypeHandler.bind(this);
    }

    componentDidMount(){
        WarrantyService.getWarrantyById(this.state.id).then( (res) =>{
            let warranty = res.data;
            this.setState({
                coverageMonths: warranty.coverageMonths,
                warrantyType: warranty.warrantyType
            });
        });
    }

    updateWarranty = (e) => {
        e.preventDefault();
        let warranty = {
            warrantyId: this.state.id,
            coverageMonths: this.state.coverageMonths,
            warrantyType: this.state.warrantyType
        };
        console.log('warranty => ' + JSON.stringify(warranty));
        console.log('id => ' + JSON.stringify(this.state.id));
        WarrantyService.updateWarranty(warranty).then( res => {
            this.props.history.push('/warrantys');
        });
    }

    changecoverageMonthsHandler= (event) => {
        this.setState({coverageMonths: event.target.value});
    }
    changeWarrantyTypeHandler= (event) => {
        this.setState({warrantyType: event.target.value});
    }

    cancel(){
        this.props.history.push('/warrantys');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Warranty</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> coverageMonths: </label>
                                                <input type="number" placeholder="coverageMonths" name="coverageMonths" className="form-control" value={this.state.coverageMonths} onChange={this.changecoverageMonthsHandler}/>

                                            <label> WarrantyType: </label>
                                                <select value={this.state.warrantyType} onChange={this.changeWarrantyTypeHandler}>
                      <option name="WarrantyType" className="form-control" >
                          Basic
                      </option>
                      <option name="WarrantyType" className="form-control" >
                          Powerplant
                      </option>
                      <option name="WarrantyType" className="form-control" >
                          Avionics
                      </option>
                      <option name="WarrantyType" className="form-control" >
                          Corrosion
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateWarranty}>Save</button>
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

export default UpdateWarrantyComponent
