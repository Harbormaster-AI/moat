
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateComponent_Component } from './create.component';
import { Component_Service } from '../../../services/Component_.service';
import { Router } from '@angular/router';

describe('CreateComponent_Component', () => {
  let component: CreateComponent_Component;
  let fixture: ComponentFixture<CreateComponent_Component>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateComponent_Component
      ],
      providers: [
        Component_Service,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateComponent_Component);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});