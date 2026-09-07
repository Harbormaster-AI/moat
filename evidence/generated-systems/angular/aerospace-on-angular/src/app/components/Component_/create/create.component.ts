import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Component_Service } from '../../../services/Component_.service';
import { Component_ } from '../../../models/Component_';
import { SubBaseComponent } from '../../Component_/sub.base.component';

@Component({
    selector: 'app-create-component_',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateComponent_Component extends SubBaseComponent implements OnInit {

    title = 'Add Component_';

    component_Form: FormGroup;
    component_: Component_;

    constructor( http: HttpClient,
        private component_Service: Component_Service,
        private fb: FormBuilder,
        private router: Router
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

    
    addComponent_(partNumber, name, Supplier, ComponentCategory, SerializationMethod): void {
        this.component_Service
        .addComponent_(partNumber, name, Supplier, ComponentCategory, SerializationMethod)
            .subscribe(() => {
                this.router.navigate(['/indexComponent_']);
            });
    }

    ngOnInit(): void {
    }
}