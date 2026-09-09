import React, { Component } from 'react'
import CostCenterService from '../services/CostCenterService';

class UpdateCostCenterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                code: '',
                name: ''
        }
        this.updateCostCenter = this.updateCostCenter.bind(this);

        this.changecodeHandler = this.changecodeHandler.bind(this);
        this.changenameHandler = this.changenameHandler.bind(this);
    }

    componentDidMount(){
        CostCenterService.getCostCenterById(this.state.id).then( (res) =>{
            let costCenter = res.data;
            this.setState({
                code: costCenter.code,
                name: costCenter.name
            });
        });
    }

    updateCostCenter = (e) => {
        e.preventDefault();
        let costCenter = {
            costCenterId: this.state.id,
            code: this.state.code,
            name: this.state.name
        };
        console.log('costCenter => ' + JSON.stringify(costCenter));
        console.log('id => ' + JSON.stringify(this.state.id));
        CostCenterService.updateCostCenter(costCenter).then( res => {
            this.props.history.push('/costCenters');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update CostCenter</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> code: </label>
                                                <input placeholder="code" name="code" className="form-control" value={this.state.code} onChange={this.changecodeHandler}/>

                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateCostCenter}>Save</button>
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

export default UpdateCostCenterComponent
