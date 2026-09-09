import React, { Component } from 'react'
import UnderwriterService from '../services/UnderwriterService';

class CreateUnderwriterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                firstName: '',
                lastName: '',
                employeeId: '',
                authorityLimit: ''
        }
        this.changefirstNameHandler = this.changefirstNameHandler.bind(this);
        this.changelastNameHandler = this.changelastNameHandler.bind(this);
        this.changeemployeeIdHandler = this.changeemployeeIdHandler.bind(this);
        this.changeauthorityLimitHandler = this.changeauthorityLimitHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            UnderwriterService.getUnderwriterById(this.state.id).then( (res) =>{
                let underwriter = res.data;
                this.setState({
                    firstName: underwriter.firstName,
                    lastName: underwriter.lastName,
                    employeeId: underwriter.employeeId,
                    authorityLimit: underwriter.authorityLimit
                });
            });
        }        
    }
    saveOrUpdateUnderwriter = (e) => {
        e.preventDefault();
        let underwriter = {
                underwriterId: this.state.id,
                firstName: this.state.firstName,
                lastName: this.state.lastName,
                employeeId: this.state.employeeId,
                authorityLimit: this.state.authorityLimit
            };
        console.log('underwriter => ' + JSON.stringify(underwriter));

        // step 5
        if(this.state.id === '_add'){
            underwriter.underwriterId=''
            UnderwriterService.createUnderwriter(underwriter).then(res =>{
                this.props.history.push('/underwriters');
            });
        }else{
            UnderwriterService.updateUnderwriter(underwriter).then( res => {
                this.props.history.push('/underwriters');
            });
        }
    }
    
    changefirstNameHandler= (event) => {
        this.setState({firstName: event.target.value});
    }
    changelastNameHandler= (event) => {
        this.setState({lastName: event.target.value});
    }
    changeemployeeIdHandler= (event) => {
        this.setState({employeeId: event.target.value});
    }
    changeauthorityLimitHandler= (event) => {
        this.setState({authorityLimit: event.target.value});
    }

    cancel(){
        this.props.history.push('/underwriters');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Underwriter</h3>
        }else{
            return <h3 className="text-center">Update Underwriter</h3>
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

                                            <label> employeeId:&emsp; </label>
                                                <input placeholder="employeeId" name="employeeId" className="form-control" value={this.state.employeeId} onChange={this.changeemployeeIdHandler}/>

                                            <label> authorityLimit:&emsp; </label>
                                                <input placeholder="authorityLimit" name="authorityLimit" className="form-control" value={this.state.authorityLimit} onChange={this.changeauthorityLimitHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateUnderwriter}>Save</button>
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

export default CreateUnderwriterComponent
