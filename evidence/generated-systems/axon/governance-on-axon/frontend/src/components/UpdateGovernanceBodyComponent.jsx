import React, { Component } from 'react'
import GovernanceBodyService from '../services/GovernanceBodyService';

class UpdateGovernanceBodyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                charterUrl: '',
                chair: '',
                bodyType: ''
        }
        this.updateGovernanceBody = this.updateGovernanceBody.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecharterUrlHandler = this.changecharterUrlHandler.bind(this);
        this.changechairHandler = this.changechairHandler.bind(this);
        this.changeBodyTypeHandler = this.changeBodyTypeHandler.bind(this);
    }

    componentDidMount(){
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

    updateGovernanceBody = (e) => {
        e.preventDefault();
        let governanceBody = {
            governanceBodyId: this.state.id,
            name: this.state.name,
            charterUrl: this.state.charterUrl,
            chair: this.state.chair,
            bodyType: this.state.bodyType
        };
        console.log('governanceBody => ' + JSON.stringify(governanceBody));
        console.log('id => ' + JSON.stringify(this.state.id));
        GovernanceBodyService.updateGovernanceBody(governanceBody).then( res => {
            this.props.history.push('/governanceBodys');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update GovernanceBody</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> charterUrl: </label>
                                                <input placeholder="charterUrl" name="charterUrl" className="form-control" value={this.state.charterUrl} onChange={this.changecharterUrlHandler}/>

                                            <label> chair: </label>
                                                <input placeholder="chair" name="chair" className="form-control" value={this.state.chair} onChange={this.changechairHandler}/>

                                            <label> BodyType: </label>
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
                                        <button className="btn btn-success" onClick={this.updateGovernanceBody}>Save</button>
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

export default UpdateGovernanceBodyComponent
