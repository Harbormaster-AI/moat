import React, { Component } from 'react'
import BusinessUnitService from '../services/BusinessUnitService';

class CreateBusinessUnitComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                leader: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeleaderHandler = this.changeleaderHandler.bind(this);
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
                    leader: businessUnit.leader
                });
            });
        }        
    }
    saveOrUpdateBusinessUnit = (e) => {
        e.preventDefault();
        let businessUnit = {
                businessUnitId: this.state.id,
                name: this.state.name,
                leader: this.state.leader
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
    changeleaderHandler= (event) => {
        this.setState({leader: event.target.value});
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

                                            <label> leader:&emsp; </label>
                                                <input placeholder="leader" name="leader" className="form-control" value={this.state.leader} onChange={this.changeleaderHandler}/>

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
