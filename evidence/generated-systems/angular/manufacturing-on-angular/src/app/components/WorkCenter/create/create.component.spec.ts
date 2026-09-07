
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateWorkCenterComponent } from './create.component';
import { WorkCenterService } from '../../../services/WorkCenter.service';
import { Router } from '@angular/router';

describe('CreateWorkCenterComponent', () => {
  let component: CreateWorkCenterComponent;
  let fixture: ComponentFixture<CreateWorkCenterComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateWorkCenterComponent
      ],
      providers: [
        WorkCenterService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateWorkCenterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});