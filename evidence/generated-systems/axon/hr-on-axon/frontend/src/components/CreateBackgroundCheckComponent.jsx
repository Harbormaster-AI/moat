import React, { Component } from 'react'
import BackgroundCheckService from '../services/BackgroundCheckService';

class CreateBackgroundCheckComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                checkNumber: '',
                provider: '',
                completedDate: '',
                status: ''
        }
        this.changecheckNumberHandler = this.changecheckNumberHandler.bind(this);
        this.changeproviderHandler = this.changeproviderHandler.bind(this);
        this.changecompletedDateHandler = this.changecompletedDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            BackgroundCheckService.getBackgroundCheckById(this.state.id).then( (res) =>{
                let backgroundCheck = res.data;
                this.setState({
                    checkNumber: backgroundCheck.checkNumber,
                    provider: backgroundCheck.provider,
                    completedDate: backgroundCheck.completedDate,
                    status: backgroundCheck.status
                });
            });
        }        
    }
    saveOrUpdateBackgroundCheck = (e) => {
        e.preventDefault();
        let backgroundCheck = {
                backgroundCheckId: this.state.id,
                checkNumber: this.state.checkNumber,
                provider: this.state.provider,
                completedDate: this.state.completedDate,
                status: this.state.status
            };
        console.log('backgroundCheck => ' + JSON.stringify(backgroundCheck));

        // step 5
        if(this.state.id === '_add'){
            backgroundCheck.backgroundCheckId=''
            BackgroundCheckService.createBackgroundCheck(backgroundCheck).then(res =>{
                this.props.history.push('/backgroundChecks');
            });
        }else{
            BackgroundCheckService.updateBackgroundCheck(backgroundCheck).then( res => {
                this.props.history.push('/backgroundChecks');
            });
        }
    }
    
    changecheckNumberHandler= (event) => {
        this.setState({checkNumber: event.target.value});
    }
    changeproviderHandler= (event) => {
        this.setState({provider: event.target.value});
    }
    changecompletedDateHandler= (event) => {
        this.setState({completedDate: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/backgroundChecks');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add BackgroundCheck</h3>
        }else{
            return <h3 className="text-center">Update BackgroundCheck</h3>
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
                                            <label> checkNumber:&emsp; </label>
                                                <input placeholder="checkNumber" name="checkNumber" className="form-control" value={this.state.checkNumber} onChange={this.changecheckNumberHandler}/>

                                            <label> provider:&emsp; </label>
                                                <input placeholder="provider" name="provider" className="form-control" value={this.state.provider} onChange={this.changeproviderHandler}/>

                                            <label> completedDate:&emsp; </label>
                                                <input type="date" placeholder="completedDate" name="completedDate" className="form-control" value={this.state.completedDate} onChange={this.changecompletedDateHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Ordered
                      </option>
                      <option name="Status" className="form-control" >
                          InProgress
                      </option>
                      <option name="Status" className="form-control" >
                          Clear
                      </option>
                      <option name="Status" className="form-control" >
                          Adverse
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateBackgroundCheck}>Save</button>
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

export default CreateBackgroundCheckComponent
