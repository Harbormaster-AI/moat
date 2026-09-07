
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateExposureComponent } from './create.component';
import { ExposureService } from '../../../services/Exposure.service';
import { Router } from '@angular/router';

describe('CreateExposureComponent', () => {
  let component: CreateExposureComponent;
  let fixture: ComponentFixture<CreateExposureComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateExposureComponent
      ],
      providers: [
        ExposureService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateExposureComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});