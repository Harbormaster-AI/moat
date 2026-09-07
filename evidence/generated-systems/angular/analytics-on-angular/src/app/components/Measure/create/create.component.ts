import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MeasureService } from '../../../services/Measure.service';
import { Measure } from '../../../models/Measure';
import { SubBaseComponent } from '../../Measure/sub.base.component';

@Component({
    selector: 'app-create-measure',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateMeasureComponent extends SubBaseComponent implements OnInit {

    title = 'Add Measure';

    measureForm: FormGroup;
    measure: Measure;

    constructor( http: HttpClient,
        private measureService: MeasureService,
        private fb: FormBuilder,
        private router: Router
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

    
    addMeasure(name, format, SemanticModel, Datasets, GlossaryTerms, Aggregation): void {
        this.measureService
        .addMeasure(name, format, SemanticModel, Datasets, GlossaryTerms, Aggregation)
            .subscribe(() => {
                this.router.navigate(['/indexMeasure']);
            });
    }

    ngOnInit(): void {
    }
}