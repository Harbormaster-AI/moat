import React, { Component } from 'react'
import BusinessGlossaryTermService from '../services/BusinessGlossaryTermService';

class CreateBusinessGlossaryTermComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                term: '',
                definition: '',
                steward: ''
        }
        this.changetermHandler = this.changetermHandler.bind(this);
        this.changedefinitionHandler = this.changedefinitionHandler.bind(this);
        this.changestewardHandler = this.changestewardHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            BusinessGlossaryTermService.getBusinessGlossaryTermById(this.state.id).then( (res) =>{
                let businessGlossaryTerm = res.data;
                this.setState({
                    term: businessGlossaryTerm.term,
                    definition: businessGlossaryTerm.definition,
                    steward: businessGlossaryTerm.steward
                });
            });
        }        
    }
    saveOrUpdateBusinessGlossaryTerm = (e) => {
        e.preventDefault();
        let businessGlossaryTerm = {
                businessGlossaryTermId: this.state.id,
                term: this.state.term,
                definition: this.state.definition,
                steward: this.state.steward
            };
        console.log('businessGlossaryTerm => ' + JSON.stringify(businessGlossaryTerm));

        // step 5
        if(this.state.id === '_add'){
            businessGlossaryTerm.businessGlossaryTermId=''
            BusinessGlossaryTermService.createBusinessGlossaryTerm(businessGlossaryTerm).then(res =>{
                this.props.history.push('/businessGlossaryTerms');
            });
        }else{
            BusinessGlossaryTermService.updateBusinessGlossaryTerm(businessGlossaryTerm).then( res => {
                this.props.history.push('/businessGlossaryTerms');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add BusinessGlossaryTerm</h3>
        }else{
            return <h3 className="text-center">Update BusinessGlossaryTerm</h3>
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
                                            <label> term:&emsp; </label>
                                                <input placeholder="term" name="term" className="form-control" value={this.state.term} onChange={this.changetermHandler}/>

                                            <label> definition:&emsp; </label>
                                                <input placeholder="definition" name="definition" className="form-control" value={this.state.definition} onChange={this.changedefinitionHandler}/>

                                            <label> steward:&emsp; </label>
                                                <input placeholder="steward" name="steward" className="form-control" value={this.state.steward} onChange={this.changestewardHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateBusinessGlossaryTerm}>Save</button>
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

export default CreateBusinessGlossaryTermComponent
