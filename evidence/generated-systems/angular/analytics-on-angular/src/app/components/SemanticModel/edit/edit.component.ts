import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { SemanticModelService } from '../../../services/SemanticModel.service';
import { SubBaseComponent } from '../../SemanticModel/sub.base.component';


@Component({
    selector: 'app-edit-semanticModel',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditSemanticModelComponent extends SubBaseComponent implements OnInit {

    title = 'Edit SemanticModel';

    semanticModelForm: FormGroup;
    semanticModel: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: SemanticModelService,
        private fb: FormBuilder
) {
        super(http);
        this.semanticModelForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      version: ['', Validators.required],
      grain: ['', Validators.required],
      Datasets: ['', ],
      Metrics: ['', ],
      Dimensions: ['', ],
      Measures: ['', ],
      GlossaryTerms: ['', ]
        });
    }

    
    updateSemanticModel(name, version, grain, Datasets, Metrics, Dimensions, Measures, GlossaryTerms): void {
        this.route.params.subscribe((params) => {

                        this.service.updateSemanticModel(name, version, grain, Datasets, Metrics, Dimensions, Measures, GlossaryTerms, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexSemanticModel']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getSemanticModel(params['id']).subscribe(res => {
                this.semanticModel = res;
            });
        });
    }
}