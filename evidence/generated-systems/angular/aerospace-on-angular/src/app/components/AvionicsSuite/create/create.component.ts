import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AvionicsSuiteService } from '../../../services/AvionicsSuite.service';
import { AvionicsSuite } from '../../../models/AvionicsSuite';
import { SubBaseComponent } from '../../AvionicsSuite/sub.base.component';

@Component({
    selector: 'app-create-avionicsSuite',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAvionicsSuiteComponent extends SubBaseComponent implements OnInit {

    title = 'Add AvionicsSuite';

    avionicsSuiteForm: FormGroup;
    avionicsSuite: AvionicsSuite;

    constructor( http: HttpClient,
        private avionicsSuiteService: AvionicsSuiteService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.avionicsSuiteForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  suiteName: ['', Validators.required],
      softwareBaseline: ['', Validators.required],
      Supplier: ['', ],
      Variants: ['', ],
      SoftwareLoads: ['', ]
        });
    }

    
    addAvionicsSuite(suiteName, softwareBaseline, Supplier, Variants, SoftwareLoads): void {
        this.avionicsSuiteService
        .addAvionicsSuite(suiteName, softwareBaseline, Supplier, Variants, SoftwareLoads)
            .subscribe(() => {
                this.router.navigate(['/indexAvionicsSuite']);
            });
    }

    ngOnInit(): void {
    }
}