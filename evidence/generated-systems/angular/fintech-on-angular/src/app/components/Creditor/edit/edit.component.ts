import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { CreditorService } from '../../../services/Creditor.service';
import { SubBaseComponent } from '../../Creditor/sub.base.component';


@Component({
    selector: 'app-edit-creditor',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditCreditorComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Creditor';

    creditorForm: FormGroup;
    creditor: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: CreditorService,
        private fb: FormBuilder
) {
        super(http);
        this.creditorForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      bic: ['', Validators.required],
      address: ['', Validators.required],
      Mandates: ['', ]
        });
    }

    
    updateCreditor(name, bic, address, Mandates): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCreditor(name, bic, address, Mandates, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexCreditor']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getCreditor(params['id']).subscribe(res => {
                this.creditor = res;
            });
        });
    }
}