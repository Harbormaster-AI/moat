import React, { Component } from 'react'
import JobFamilyService from '../services/JobFamilyService';

class CreateJobFamilyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                description: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            JobFamilyService.getJobFamilyById(this.state.id).then( (res) =>{
                let jobFamily = res.data;
                this.setState({
                    name: jobFamily.name,
                    description: jobFamily.description
                });
            });
        }        
    }
    saveOrUpdateJobFamily = (e) => {
        e.preventDefault();
        let jobFamily = {
                jobFamilyId: this.state.id,
                name: this.state.name,
                description: this.state.description
            };
        console.log('jobFamily => ' + JSON.stringify(jobFamily));

        // step 5
        if(this.state.id === '_add'){
            jobFamily.jobFamilyId=''
            JobFamilyService.createJobFamily(jobFamily).then(res =>{
                this.props.history.push('/jobFamilys');
            });
        }else{
            JobFamilyService.updateJobFamily(jobFamily).then( res => {
                this.props.history.push('/jobFamilys');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changedescriptionHandler= (event) => {
        this.setState({description: event.target.value});
    }

    cancel(){
        this.props.history.push('/jobFamilys');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add JobFamily</h3>
        }else{
            return <h3 className="text-center">Update JobFamily</h3>
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

                                            <label> description:&emsp; </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateJobFamily}>Save</button>
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

export default CreateJobFamilyComponent
