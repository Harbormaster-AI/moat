import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AvionicsSuiteService } from '../../../services/AvionicsSuite.service';
import { SubBaseComponent } from '../../AvionicsSuite/sub.base.component';


@Component({
    selector: 'app-edit-avionicsSuite',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAvionicsSuiteComponent extends SubBaseComponent implements OnInit {

    title = 'Edit AvionicsSuite';

    avionicsSuiteForm: FormGroup;
    avionicsSuite: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AvionicsSuiteService,
        private fb: FormBuilder
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

    
    updateAvionicsSuite(suiteName, softwareBaseline, Supplier, Variants, SoftwareLoads): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAvionicsSuite(suiteName, softwareBaseline, Supplier, Variants, SoftwareLoads, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAvionicsSuite']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAvionicsSuite(params['id']).subscribe(res => {
                this.avionicsSuite = res;
            });
        });
    }
}