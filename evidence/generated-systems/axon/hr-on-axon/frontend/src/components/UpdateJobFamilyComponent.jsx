import React, { Component } from 'react'
import JobFamilyService from '../services/JobFamilyService';

class UpdateJobFamilyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                description: ''
        }
        this.updateJobFamily = this.updateJobFamily.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
    }

    componentDidMount(){
        JobFamilyService.getJobFamilyById(this.state.id).then( (res) =>{
            let jobFamily = res.data;
            this.setState({
                name: jobFamily.name,
                description: jobFamily.description
            });
        });
    }

    updateJobFamily = (e) => {
        e.preventDefault();
        let jobFamily = {
            jobFamilyId: this.state.id,
            name: this.state.name,
            description: this.state.description
        };
        console.log('jobFamily => ' + JSON.stringify(jobFamily));
        console.log('id => ' + JSON.stringify(this.state.id));
        JobFamilyService.updateJobFamily(jobFamily).then( res => {
            this.props.history.push('/jobFamilys');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update JobFamily</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> description: </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateJobFamily}>Save</button>
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

export default UpdateJobFamilyComponent
