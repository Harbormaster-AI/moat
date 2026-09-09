import React, { Component } from 'react'
import AdjusterService from '../services/AdjusterService';

class UpdateAdjusterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                firstName: '',
                lastName: '',
                licenseNumber: '',
                adjusterType: ''
        }
        this.updateAdjuster = this.updateAdjuster.bind(this);

        this.changefirstNameHandler = this.changefirstNameHandler.bind(this);
        this.changelastNameHandler = this.changelastNameHandler.bind(this);
        this.changelicenseNumberHandler = this.changelicenseNumberHandler.bind(this);
        this.changeAdjusterTypeHandler = this.changeAdjusterTypeHandler.bind(this);
    }

    componentDidMount(){
        AdjusterService.getAdjusterById(this.state.id).then( (res) =>{
            let adjuster = res.data;
            this.setState({
                firstName: adjuster.firstName,
                lastName: adjuster.lastName,
                licenseNumber: adjuster.licenseNumber,
                adjusterType: adjuster.adjusterType
            });
        });
    }

    updateAdjuster = (e) => {
        e.preventDefault();
        let adjuster = {
            adjusterId: this.state.id,
            firstName: this.state.firstName,
            lastName: this.state.lastName,
            licenseNumber: this.state.licenseNumber,
            adjusterType: this.state.adjusterType
        };
        console.log('adjuster => ' + JSON.stringify(adjuster));
        console.log('id => ' + JSON.stringify(this.state.id));
        AdjusterService.updateAdjuster(adjuster).then( res => {
            this.props.history.push('/adjusters');
        });
    }

    changefirstNameHandler= (event) => {
        this.setState({firstName: event.target.value});
    }
    changelastNameHandler= (event) => {
        this.setState({lastName: event.target.value});
    }
    changelicenseNumberHandler= (event) => {
        this.setState({licenseNumber: event.target.value});
    }
    changeAdjusterTypeHandler= (event) => {
        this.setState({adjusterType: event.target.value});
    }

    cancel(){
        this.props.history.push('/adjusters');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Adjuster</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> firstName: </label>
                                                <input placeholder="firstName" name="firstName" className="form-control" value={this.state.firstName} onChange={this.changefirstNameHandler}/>

                                            <label> lastName: </label>
                                                <input placeholder="lastName" name="lastName" className="form-control" value={this.state.lastName} onChange={this.changelastNameHandler}/>

                                            <label> licenseNumber: </label>
                                                <input placeholder="licenseNumber" name="licenseNumber" className="form-control" value={this.state.licenseNumber} onChange={this.changelicenseNumberHandler}/>

                                            <label> AdjusterType: </label>
                                                <select value={this.state.adjusterType} onChange={this.changeAdjusterTypeHandler}>
                      <option name="AdjusterType" className="form-control" >
                          Staff
                      </option>
                      <option name="AdjusterType" className="form-control" >
                          Independent
                      </option>
                      <option name="AdjusterType" className="form-control" >
                          Public
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateAdjuster}>Save</button>
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

export default UpdateAdjusterComponent
