
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateLaboratoryComponent } from './create.component';
import { LaboratoryService } from '../../../services/Laboratory.service';
import { Router } from '@angular/router';

describe('CreateLaboratoryComponent', () => {
  let component: CreateLaboratoryComponent;
  let fixture: ComponentFixture<CreateLaboratoryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateLaboratoryComponent
      ],
      providers: [
        LaboratoryService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateLaboratoryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});