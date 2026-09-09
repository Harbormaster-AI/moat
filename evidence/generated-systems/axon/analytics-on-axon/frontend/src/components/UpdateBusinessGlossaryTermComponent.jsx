import React, { Component } from 'react'
import BusinessGlossaryTermService from '../services/BusinessGlossaryTermService';

class UpdateBusinessGlossaryTermComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                term: '',
                definition: '',
                steward: ''
        }
        this.updateBusinessGlossaryTerm = this.updateBusinessGlossaryTerm.bind(this);

        this.changetermHandler = this.changetermHandler.bind(this);
        this.changedefinitionHandler = this.changedefinitionHandler.bind(this);
        this.changestewardHandler = this.changestewardHandler.bind(this);
    }

    componentDidMount(){
        BusinessGlossaryTermService.getBusinessGlossaryTermById(this.state.id).then( (res) =>{
            let businessGlossaryTerm = res.data;
            this.setState({
                term: businessGlossaryTerm.term,
                definition: businessGlossaryTerm.definition,
                steward: businessGlossaryTerm.steward
            });
        });
    }

    updateBusinessGlossaryTerm = (e) => {
        e.preventDefault();
        let businessGlossaryTerm = {
            businessGlossaryTermId: this.state.id,
            term: this.state.term,
            definition: this.state.definition,
            steward: this.state.steward
        };
        console.log('businessGlossaryTerm => ' + JSON.stringify(businessGlossaryTerm));
        console.log('id => ' + JSON.stringify(this.state.id));
        BusinessGlossaryTermService.updateBusinessGlossaryTerm(businessGlossaryTerm).then( res => {
            this.props.history.push('/businessGlossaryTerms');
        });
    }

    changetermHandler= (event) => {
        this.setState({term: event.target.value});
    }
    changedefinitionHandler= (event) => {
        this.setState({definition: event.target.value});
    }
    changestewardHandler= (event) => {
        this.setState({steward: event.target.value});
    }

    cancel(){
        this.props.history.push('/businessGlossaryTerms');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update BusinessGlossaryTerm</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> term: </label>
                                                <input placeholder="term" name="term" className="form-control" value={this.state.term} onChange={this.changetermHandler}/>

                                            <label> definition: </label>
                                                <input placeholder="definition" name="definition" className="form-control" value={this.state.definition} onChange={this.changedefinitionHandler}/>

                                            <label> steward: </label>
                                                <input placeholder="steward" name="steward" className="form-control" value={this.state.steward} onChange={this.changestewardHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateBusinessGlossaryTerm}>Save</button>
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

export default UpdateBusinessGlossaryTermComponent
