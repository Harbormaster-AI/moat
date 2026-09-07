
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateFeatureSetComponent } from './create.component';
import { FeatureSetService } from '../../../services/FeatureSet.service';
import { Router } from '@angular/router';

describe('CreateFeatureSetComponent', () => {
  let component: CreateFeatureSetComponent;
  let fixture: ComponentFixture<CreateFeatureSetComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateFeatureSetComponent
      ],
      providers: [
        FeatureSetService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateFeatureSetComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});