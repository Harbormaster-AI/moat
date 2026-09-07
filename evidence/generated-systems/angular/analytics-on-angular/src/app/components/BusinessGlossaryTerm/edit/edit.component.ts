import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { BusinessGlossaryTermService } from '../../../services/BusinessGlossaryTerm.service';
import { SubBaseComponent } from '../../BusinessGlossaryTerm/sub.base.component';


@Component({
    selector: 'app-edit-businessGlossaryTerm',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditBusinessGlossaryTermComponent extends SubBaseComponent implements OnInit {

    title = 'Edit BusinessGlossaryTerm';

    businessGlossaryTermForm: FormGroup;
    businessGlossaryTerm: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: BusinessGlossaryTermService,
        private fb: FormBuilder
) {
        super(http);
        this.businessGlossaryTermForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  term: ['', Validators.required],
      definition: ['', Validators.required],
      steward: ['', Validators.required],
      RelatedTerms: ['', ],
      Metrics: ['', ],
      Datasets: ['', ],
      Dimensions: ['', ],
      Measures: ['', ]
        });
    }

    
    updateBusinessGlossaryTerm(term, definition, steward, RelatedTerms, Metrics, Datasets, Dimensions, Measures): void {
        this.route.params.subscribe((params) => {

                        this.service.updateBusinessGlossaryTerm(term, definition, steward, RelatedTerms, Metrics, Datasets, Dimensions, Measures, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexBusinessGlossaryTerm']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getBusinessGlossaryTerm(params['id']).subscribe(res => {
                this.businessGlossaryTerm = res;
            });
        });
    }
}