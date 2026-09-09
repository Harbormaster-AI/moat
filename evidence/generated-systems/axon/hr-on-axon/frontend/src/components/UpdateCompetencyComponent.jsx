import React, { Component } from 'react'
import CompetencyService from '../services/CompetencyService';

class UpdateCompetencyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                category: ''
        }
        this.updateCompetency = this.updateCompetency.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecategoryHandler = this.changecategoryHandler.bind(this);
    }

    componentDidMount(){
        CompetencyService.getCompetencyById(this.state.id).then( (res) =>{
            let competency = res.data;
            this.setState({
                name: competency.name,
                category: competency.category
            });
        });
    }

    updateCompetency = (e) => {
        e.preventDefault();
        let competency = {
            competencyId: this.state.id,
            name: this.state.name,
            category: this.state.category
        };
        console.log('competency => ' + JSON.stringify(competency));
        console.log('id => ' + JSON.stringify(this.state.id));
        CompetencyService.updateCompetency(competency).then( res => {
            this.props.history.push('/competencys');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changecategoryHandler= (event) => {
        this.setState({category: event.target.value});
    }

    cancel(){
        this.props.history.push('/competencys');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Competency</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> category: </label>
                                                <input placeholder="category" name="category" className="form-control" value={this.state.category} onChange={this.changecategoryHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateCompetency}>Save</button>
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

export default UpdateCompetencyComponent
