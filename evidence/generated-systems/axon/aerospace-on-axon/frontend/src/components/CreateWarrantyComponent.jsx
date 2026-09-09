import React, { Component } from 'react'
import WarrantyService from '../services/WarrantyService';

class CreateWarrantyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                coverageMonths: '',
                warrantyType: ''
        }
        this.changecoverageMonthsHandler = this.changecoverageMonthsHandler.bind(this);
        this.changeWarrantyTypeHandler = this.changeWarrantyTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            WarrantyService.getWarrantyById(this.state.id).then( (res) =>{
                let warranty = res.data;
                this.setState({
                    coverageMonths: warranty.coverageMonths,
                    warrantyType: warranty.warrantyType
                });
            });
        }        
    }
    saveOrUpdateWarranty = (e) => {
        e.preventDefault();
        let warranty = {
                warrantyId: this.state.id,
                coverageMonths: this.state.coverageMonths,
                warrantyType: this.state.warrantyType
            };
        console.log('warranty => ' + JSON.stringify(warranty));

        // step 5
        if(this.state.id === '_add'){
            warranty.warrantyId=''
            WarrantyService.createWarranty(warranty).then(res =>{
                this.props.history.push('/warrantys');
            });
        }else{
            WarrantyService.updateWarranty(warranty).then( res => {
                this.props.history.push('/warrantys');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Warranty</h3>
        }else{
            return <h3 className="text-center">Update Warranty</h3>
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
                                            <label> coverageMonths:&emsp; </label>
                                                <input type="number" placeholder="coverageMonths" name="coverageMonths" className="form-control" value={this.state.coverageMonths} onChange={this.changecoverageMonthsHandler}/>

                                            <label> WarrantyType:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateWarranty}>Save</button>
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

export default CreateWarrantyComponent
