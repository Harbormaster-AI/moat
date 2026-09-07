
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateException_Component } from './create.component';
import { Exception_Service } from '../../../services/Exception_.service';
import { Router } from '@angular/router';

describe('CreateException_Component', () => {
  let component: CreateException_Component;
  let fixture: ComponentFixture<CreateException_Component>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateException_Component
      ],
      providers: [
        Exception_Service,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateException_Component);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});