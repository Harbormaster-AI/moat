
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateModel_Component } from './create.component';
import { Model_Service } from '../../../services/Model_.service';
import { Router } from '@angular/router';

describe('CreateModel_Component', () => {
  let component: CreateModel_Component;
  let fixture: ComponentFixture<CreateModel_Component>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateModel_Component
      ],
      providers: [
        Model_Service,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateModel_Component);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});