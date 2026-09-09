import React, { Component } from 'react'
import GovernanceBodyService from '../services/GovernanceBodyService';

class CreateGovernanceBodyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                charterUrl: '',
                chair: '',
                bodyType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecharterUrlHandler = this.changecharterUrlHandler.bind(this);
        this.changechairHandler = this.changechairHandler.bind(this);
        this.changeBodyTypeHandler = this.changeBodyTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            GovernanceBodyService.getGovernanceBodyById(this.state.id).then( (res) =>{
                let governanceBody = res.data;
                this.setState({
                    name: governanceBody.name,
                    charterUrl: governanceBody.charterUrl,
                    chair: governanceBody.chair,
                    bodyType: governanceBody.bodyType
                });
            });
        }        
    }
    saveOrUpdateGovernanceBody = (e) => {
        e.preventDefault();
        let governanceBody = {
                governanceBodyId: this.state.id,
                name: this.state.name,
                charterUrl: this.state.charterUrl,
                chair: this.state.chair,
                bodyType: this.state.bodyType
            };
        console.log('governanceBody => ' + JSON.stringify(governanceBody));

        // step 5
        if(this.state.id === '_add'){
            governanceBody.governanceBodyId=''
            GovernanceBodyService.createGovernanceBody(governanceBody).then(res =>{
                this.props.history.push('/governanceBodys');
            });
        }else{
            GovernanceBodyService.updateGovernanceBody(governanceBody).then( res => {
                this.props.history.push('/governanceBodys');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changecharterUrlHandler= (event) => {
        this.setState({charterUrl: event.target.value});
    }
    changechairHandler= (event) => {
        this.setState({chair: event.target.value});
    }
    changeBodyTypeHandler= (event) => {
        this.setState({bodyType: event.target.value});
    }

    cancel(){
        this.props.history.push('/governanceBodys');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add GovernanceBody</h3>
        }else{
            return <h3 className="text-center">Update GovernanceBody</h3>
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

                                            <label> charterUrl:&emsp; </label>
                                                <input placeholder="charterUrl" name="charterUrl" className="form-control" value={this.state.charterUrl} onChange={this.changecharterUrlHandler}/>

                                            <label> chair:&emsp; </label>
                                                <input placeholder="chair" name="chair" className="form-control" value={this.state.chair} onChange={this.changechairHandler}/>

                                            <label> BodyType:&emsp; </label>
                                                <select value={this.state.bodyType} onChange={this.changeBodyTypeHandler}>
                      <option name="BodyType" className="form-control" >
                          Board
                      </option>
                      <option name="BodyType" className="form-control" >
                          Committee
                      </option>
                      <option name="BodyType" className="form-control" >
                          Council
                      </option>
                      <option name="BodyType" className="form-control" >
                          WorkingGroup
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateGovernanceBody}>Save</button>
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

export default CreateGovernanceBodyComponent
