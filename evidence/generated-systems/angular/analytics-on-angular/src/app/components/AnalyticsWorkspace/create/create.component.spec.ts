
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAnalyticsWorkspaceComponent } from './create.component';
import { AnalyticsWorkspaceService } from '../../../services/AnalyticsWorkspace.service';
import { Router } from '@angular/router';

describe('CreateAnalyticsWorkspaceComponent', () => {
  let component: CreateAnalyticsWorkspaceComponent;
  let fixture: ComponentFixture<CreateAnalyticsWorkspaceComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateAnalyticsWorkspaceComponent
      ],
      providers: [
        AnalyticsWorkspaceService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAnalyticsWorkspaceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});