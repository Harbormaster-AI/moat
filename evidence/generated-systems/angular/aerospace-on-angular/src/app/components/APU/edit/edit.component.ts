import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { APUService } from '../../../services/APU.service';
import { SubBaseComponent } from '../../APU/sub.base.component';


@Component({
    selector: 'app-edit-aPU',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAPUComponent extends SubBaseComponent implements OnInit {

    title = 'Edit APU';

    aPUForm: FormGroup;
    aPU: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: APUService,
        private fb: FormBuilder
) {
        super(http);
        this.aPUForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  model_: ['', Validators.required],
      Supplier: ['', ],
      Variants: ['', ]
        });
    }

    
    updateAPU(model_, Supplier, Variants): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAPU(model_, Supplier, Variants, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAPU']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAPU(params['id']).subscribe(res => {
                this.aPU = res;
            });
        });
    }
}