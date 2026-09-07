import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { MeasureService } from '../../../services/Measure.service';
import { SubBaseComponent } from '../../Measure/sub.base.component';


@Component({
    selector: 'app-edit-measure',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditMeasureComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Measure';

    measureForm: FormGroup;
    measure: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: MeasureService,
        private fb: FormBuilder
) {
        super(http);
        this.measureForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      format: ['', Validators.required],
      SemanticModel: ['', ],
      Datasets: ['', ],
      GlossaryTerms: ['', ],
      Aggregation: ['', ]
        });
    }

    
    updateMeasure(name, format, SemanticModel, Datasets, GlossaryTerms, Aggregation): void {
        this.route.params.subscribe((params) => {

                        this.service.updateMeasure(name, format, SemanticModel, Datasets, GlossaryTerms, Aggregation, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexMeasure']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getMeasure(params['id']).subscribe(res => {
                this.measure = res;
            });
        });
    }
}