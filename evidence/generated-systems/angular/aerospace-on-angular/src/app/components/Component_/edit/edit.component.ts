import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { Component_Service } from '../../../services/Component_.service';
import { SubBaseComponent } from '../../Component_/sub.base.component';


@Component({
    selector: 'app-edit-component_',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditComponent_Component extends SubBaseComponent implements OnInit {

    title = 'Edit Component_';

    component_Form: FormGroup;
    component_: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: Component_Service,
        private fb: FormBuilder
) {
        super(http);
        this.component_Form = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  partNumber: ['', Validators.required],
      name: ['', Validators.required],
      Supplier: ['', ],
      ComponentCategory: ['', ],
      SerializationMethod: ['', ]
        });
    }

    
    updateComponent_(partNumber, name, Supplier, ComponentCategory, SerializationMethod): void {
        this.route.params.subscribe((params) => {

                        this.service.updateComponent_(partNumber, name, Supplier, ComponentCategory, SerializationMethod, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexComponent_']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getComponent_(params['id']).subscribe(res => {
                this.component_ = res;
            });
        });
    }
}