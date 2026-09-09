import React, { Component } from 'react'
import CompetencyService from '../services/CompetencyService';

class CreateCompetencyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                category: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecategoryHandler = this.changecategoryHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            CompetencyService.getCompetencyById(this.state.id).then( (res) =>{
                let competency = res.data;
                this.setState({
                    name: competency.name,
                    category: competency.category
                });
            });
        }        
    }
    saveOrUpdateCompetency = (e) => {
        e.preventDefault();
        let competency = {
                competencyId: this.state.id,
                name: this.state.name,
                category: this.state.category
            };
        console.log('competency => ' + JSON.stringify(competency));

        // step 5
        if(this.state.id === '_add'){
            competency.competencyId=''
            CompetencyService.createCompetency(competency).then(res =>{
                this.props.history.push('/competencys');
            });
        }else{
            CompetencyService.updateCompetency(competency).then( res => {
                this.props.history.push('/competencys');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Competency</h3>
        }else{
            return <h3 className="text-center">Update Competency</h3>
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

                                            <label> category:&emsp; </label>
                                                <input placeholder="category" name="category" className="form-control" value={this.state.category} onChange={this.changecategoryHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateCompetency}>Save</button>
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

export default CreateCompetencyComponent
