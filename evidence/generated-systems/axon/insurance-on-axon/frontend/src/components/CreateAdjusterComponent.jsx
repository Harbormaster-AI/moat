import React, { Component } from 'react'
import AdjusterService from '../services/AdjusterService';

class CreateAdjusterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                firstName: '',
                lastName: '',
                licenseNumber: '',
                adjusterType: ''
        }
        this.changefirstNameHandler = this.changefirstNameHandler.bind(this);
        this.changelastNameHandler = this.changelastNameHandler.bind(this);
        this.changelicenseNumberHandler = this.changelicenseNumberHandler.bind(this);
        this.changeAdjusterTypeHandler = this.changeAdjusterTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateAdjuster = (e) => {
        e.preventDefault();
        let adjuster = {
                adjusterId: this.state.id,
                firstName: this.state.firstName,
                lastName: this.state.lastName,
                licenseNumber: this.state.licenseNumber,
                adjusterType: this.state.adjusterType
            };
        console.log('adjuster => ' + JSON.stringify(adjuster));

        // step 5
        if(this.state.id === '_add'){
            adjuster.adjusterId=''
            AdjusterService.createAdjuster(adjuster).then(res =>{
                this.props.history.push('/adjusters');
            });
        }else{
            AdjusterService.updateAdjuster(adjuster).then( res => {
                this.props.history.push('/adjusters');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Adjuster</h3>
        }else{
            return <h3 className="text-center">Update Adjuster</h3>
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
                                            <label> firstName:&emsp; </label>
                                                <input placeholder="firstName" name="firstName" className="form-control" value={this.state.firstName} onChange={this.changefirstNameHandler}/>

                                            <label> lastName:&emsp; </label>
                                                <input placeholder="lastName" name="lastName" className="form-control" value={this.state.lastName} onChange={this.changelastNameHandler}/>

                                            <label> licenseNumber:&emsp; </label>
                                                <input placeholder="licenseNumber" name="licenseNumber" className="form-control" value={this.state.licenseNumber} onChange={this.changelicenseNumberHandler}/>

                                            <label> AdjusterType:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAdjuster}>Save</button>
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

export default CreateAdjusterComponent
