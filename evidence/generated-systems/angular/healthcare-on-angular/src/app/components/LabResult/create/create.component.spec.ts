
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateLabResultComponent } from './create.component';
import { LabResultService } from '../../../services/LabResult.service';
import { Router } from '@angular/router';

describe('CreateLabResultComponent', () => {
  let component: CreateLabResultComponent;
  let fixture: ComponentFixture<CreateLabResultComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateLabResultComponent
      ],
      providers: [
        LabResultService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateLabResultComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});