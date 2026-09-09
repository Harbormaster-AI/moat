import React, { Component } from 'react'
import BusinessUnitService from '../services/BusinessUnitService';

class UpdateBusinessUnitComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                leader: ''
        }
        this.updateBusinessUnit = this.updateBusinessUnit.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeleaderHandler = this.changeleaderHandler.bind(this);
    }

    componentDidMount(){
        BusinessUnitService.getBusinessUnitById(this.state.id).then( (res) =>{
            let businessUnit = res.data;
            this.setState({
                name: businessUnit.name,
                leader: businessUnit.leader
            });
        });
    }

    updateBusinessUnit = (e) => {
        e.preventDefault();
        let businessUnit = {
            businessUnitId: this.state.id,
            name: this.state.name,
            leader: this.state.leader
        };
        console.log('businessUnit => ' + JSON.stringify(businessUnit));
        console.log('id => ' + JSON.stringify(this.state.id));
        BusinessUnitService.updateBusinessUnit(businessUnit).then( res => {
            this.props.history.push('/businessUnits');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update BusinessUnit</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> leader: </label>
                                                <input placeholder="leader" name="leader" className="form-control" value={this.state.leader} onChange={this.changeleaderHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateBusinessUnit}>Save</button>
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

export default UpdateBusinessUnitComponent
