
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditAnalyticsWorkspaceComponent } from './edit.component';
import { AnalyticsWorkspaceService } from '../../../services/AnalyticsWorkspace.service';

describe('EditAnalyticsWorkspaceComponent', () => {
  let component: EditAnalyticsWorkspaceComponent;
  let fixture: ComponentFixture<EditAnalyticsWorkspaceComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditAnalyticsWorkspaceComponent
      ],
      providers: [
        AnalyticsWorkspaceService,
        {
          provide: ActivatedRoute,
          useValue: {
            params: of({ id: '1' })
          }
        },
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(EditAnalyticsWorkspaceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});