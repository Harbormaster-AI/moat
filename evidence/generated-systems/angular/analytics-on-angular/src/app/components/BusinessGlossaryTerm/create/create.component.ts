import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { BusinessGlossaryTermService } from '../../../services/BusinessGlossaryTerm.service';
import { BusinessGlossaryTerm } from '../../../models/BusinessGlossaryTerm';
import { SubBaseComponent } from '../../BusinessGlossaryTerm/sub.base.component';

@Component({
    selector: 'app-create-businessGlossaryTerm',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateBusinessGlossaryTermComponent extends SubBaseComponent implements OnInit {

    title = 'Add BusinessGlossaryTerm';

    businessGlossaryTermForm: FormGroup;
    businessGlossaryTerm: BusinessGlossaryTerm;

    constructor( http: HttpClient,
        private businessGlossaryTermService: BusinessGlossaryTermService,
        private fb: FormBuilder,
        private router: Router
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

    
    addBusinessGlossaryTerm(term, definition, steward, RelatedTerms, Metrics, Datasets, Dimensions, Measures): void {
        this.businessGlossaryTermService
        .addBusinessGlossaryTerm(term, definition, steward, RelatedTerms, Metrics, Datasets, Dimensions, Measures)
            .subscribe(() => {
                this.router.navigate(['/indexBusinessGlossaryTerm']);
            });
    }

    ngOnInit(): void {
    }
}