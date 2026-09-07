
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateRecord_Component } from './create.component';
import { Record_Service } from '../../../services/Record_.service';
import { Router } from '@angular/router';

describe('CreateRecord_Component', () => {
  let component: CreateRecord_Component;
  let fixture: ComponentFixture<CreateRecord_Component>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateRecord_Component
      ],
      providers: [
        Record_Service,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateRecord_Component);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});