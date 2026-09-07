
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateDataProcessingActivityComponent } from './create.component';
import { DataProcessingActivityService } from '../../../services/DataProcessingActivity.service';
import { Router } from '@angular/router';

describe('CreateDataProcessingActivityComponent', () => {
  let component: CreateDataProcessingActivityComponent;
  let fixture: ComponentFixture<CreateDataProcessingActivityComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateDataProcessingActivityComponent
      ],
      providers: [
        DataProcessingActivityService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateDataProcessingActivityComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});