import React, { Component } from 'react'
import BusinessUnitService from '../services/BusinessUnitService';

class CreateBusinessUnitComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                code: '',
                category: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecodeHandler = this.changecodeHandler.bind(this);
        this.changeCategoryHandler = this.changeCategoryHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            BusinessUnitService.getBusinessUnitById(this.state.id).then( (res) =>{
                let businessUnit = res.data;
                this.setState({
                    name: businessUnit.name,
                    code: businessUnit.code,
                    category: businessUnit.category
                });
            });
        }        
    }
    saveOrUpdateBusinessUnit = (e) => {
        e.preventDefault();
        let businessUnit = {
                businessUnitId: this.state.id,
                name: this.state.name,
                code: this.state.code,
                category: this.state.category
            };
        console.log('businessUnit => ' + JSON.stringify(businessUnit));

        // step 5
        if(this.state.id === '_add'){
            businessUnit.businessUnitId=''
            BusinessUnitService.createBusinessUnit(businessUnit).then(res =>{
                this.props.history.push('/businessUnits');
            });
        }else{
            BusinessUnitService.updateBusinessUnit(businessUnit).then( res => {
                this.props.history.push('/businessUnits');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changecodeHandler= (event) => {
        this.setState({code: event.target.value});
    }
    changeCategoryHandler= (event) => {
        this.setState({category: event.target.value});
    }

    cancel(){
        this.props.history.push('/businessUnits');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add BusinessUnit</h3>
        }else{
            return <h3 className="text-center">Update BusinessUnit</h3>
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
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> code:&emsp; </label>
                                                <input placeholder="code" name="code" className="form-control" value={this.state.code} onChange={this.changecodeHandler}/>

                                            <label> Category:&emsp; </label>
                                                <select value={this.state.category} onChange={this.changeCategoryHandler}>
                      <option name="Category" className="form-control" >
                          ConsumerGoods
                      </option>
                      <option name="Category" className="form-control" >
                          IndustrialEquipment
                      </option>
                      <option name="Category" className="form-control" >
                          Electronics
                      </option>
                      <option name="Category" className="form-control" >
                          Pharmaceuticals
                      </option>
                      <option name="Category" className="form-control" >
                          FoodBeverage
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateBusinessUnit}>Save</button>
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

export default CreateBusinessUnitComponent
