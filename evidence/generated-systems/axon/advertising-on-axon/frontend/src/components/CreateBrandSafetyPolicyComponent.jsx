import React, { Component } from 'react'
import BrandSafetyPolicyService from '../services/BrandSafetyPolicyService';

class CreateBrandSafetyPolicyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                level: '',
                contentRatingThreshold: ''
        }
        this.changeLevelHandler = this.changeLevelHandler.bind(this);
        this.changeContentRatingThresholdHandler = this.changeContentRatingThresholdHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            BrandSafetyPolicyService.getBrandSafetyPolicyById(this.state.id).then( (res) =>{
                let brandSafetyPolicy = res.data;
                this.setState({
                    level: brandSafetyPolicy.level,
                    contentRatingThreshold: brandSafetyPolicy.contentRatingThreshold
                });
            });
        }        
    }
    saveOrUpdateBrandSafetyPolicy = (e) => {
        e.preventDefault();
        let brandSafetyPolicy = {
                brandSafetyPolicyId: this.state.id,
                level: this.state.level,
                contentRatingThreshold: this.state.contentRatingThreshold
            };
        console.log('brandSafetyPolicy => ' + JSON.stringify(brandSafetyPolicy));

        // step 5
        if(this.state.id === '_add'){
            brandSafetyPolicy.brandSafetyPolicyId=''
            BrandSafetyPolicyService.createBrandSafetyPolicy(brandSafetyPolicy).then(res =>{
                this.props.history.push('/brandSafetyPolicys');
            });
        }else{
            BrandSafetyPolicyService.updateBrandSafetyPolicy(brandSafetyPolicy).then( res => {
                this.props.history.push('/brandSafetyPolicys');
            });
        }
    }
    
    changeLevelHandler= (event) => {
        this.setState({level: event.target.value});
    }
    changeContentRatingThresholdHandler= (event) => {
        this.setState({contentRatingThreshold: event.target.value});
    }

    cancel(){
        this.props.history.push('/brandSafetyPolicys');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add BrandSafetyPolicy</h3>
        }else{
            return <h3 className="text-center">Update BrandSafetyPolicy</h3>
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
                                            <label> Level:&emsp; </label>
                                                <select value={this.state.level} onChange={this.changeLevelHandler}>
                      <option name="Level" className="form-control" >
                          None
                      </option>
                      <option name="Level" className="form-control" >
                          Moderate
                      </option>
                      <option name="Level" className="form-control" >
                          Strict
                      </option>
                    </select>

                                            <label> ContentRatingThreshold:&emsp; </label>
                                                <select value={this.state.contentRatingThreshold} onChange={this.changeContentRatingThresholdHandler}>
                      <option name="ContentRatingThreshold" className="form-control" >
                          G
                      </option>
                      <option name="ContentRatingThreshold" className="form-control" >
                          PG
                      </option>
                      <option name="ContentRatingThreshold" className="form-control" >
                          PGThirteen
                      </option>
                      <option name="ContentRatingThreshold" className="form-control" >
                          R
                      </option>
                      <option name="ContentRatingThreshold" className="form-control" >
                          Mature
                      </option>
                      <option name="ContentRatingThreshold" className="form-control" >
                          Unrated
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateBrandSafetyPolicy}>Save</button>
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

export default CreateBrandSafetyPolicyComponent
