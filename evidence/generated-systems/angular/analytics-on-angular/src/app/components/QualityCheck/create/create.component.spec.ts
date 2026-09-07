
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateQualityCheckComponent } from './create.component';
import { QualityCheckService } from '../../../services/QualityCheck.service';
import { Router } from '@angular/router';

describe('CreateQualityCheckComponent', () => {
  let component: CreateQualityCheckComponent;
  let fixture: ComponentFixture<CreateQualityCheckComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateQualityCheckComponent
      ],
      providers: [
        QualityCheckService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateQualityCheckComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});