
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateCase_Component } from './create.component';
import { Case_Service } from '../../../services/Case_.service';
import { Router } from '@angular/router';

describe('CreateCase_Component', () => {
  let component: CreateCase_Component;
  let fixture: ComponentFixture<CreateCase_Component>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateCase_Component
      ],
      providers: [
        Case_Service,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateCase_Component);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});