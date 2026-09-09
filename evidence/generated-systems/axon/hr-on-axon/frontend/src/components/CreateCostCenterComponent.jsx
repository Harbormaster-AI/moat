import React, { Component } from 'react'
import CostCenterService from '../services/CostCenterService';

class CreateCostCenterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                code: '',
                name: ''
        }
        this.changecodeHandler = this.changecodeHandler.bind(this);
        this.changenameHandler = this.changenameHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            CostCenterService.getCostCenterById(this.state.id).then( (res) =>{
                let costCenter = res.data;
                this.setState({
                    code: costCenter.code,
                    name: costCenter.name
                });
            });
        }        
    }
    saveOrUpdateCostCenter = (e) => {
        e.preventDefault();
        let costCenter = {
                costCenterId: this.state.id,
                code: this.state.code,
                name: this.state.name
            };
        console.log('costCenter => ' + JSON.stringify(costCenter));

        // step 5
        if(this.state.id === '_add'){
            costCenter.costCenterId=''
            CostCenterService.createCostCenter(costCenter).then(res =>{
                this.props.history.push('/costCenters');
            });
        }else{
            CostCenterService.updateCostCenter(costCenter).then( res => {
                this.props.history.push('/costCenters');
            });
        }
    }
    
    changecodeHandler= (event) => {
        this.setState({code: event.target.value});
    }
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }

    cancel(){
        this.props.history.push('/costCenters');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add CostCenter</h3>
        }else{
            return <h3 className="text-center">Update CostCenter</h3>
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
                                            <label> code:&emsp; </label>
                                                <input placeholder="code" name="code" className="form-control" value={this.state.code} onChange={this.changecodeHandler}/>

                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateCostCenter}>Save</button>
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

export default CreateCostCenterComponent
