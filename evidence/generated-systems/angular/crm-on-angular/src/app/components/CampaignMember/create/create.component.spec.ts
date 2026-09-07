
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateCampaignMemberComponent } from './create.component';
import { CampaignMemberService } from '../../../services/CampaignMember.service';
import { Router } from '@angular/router';

describe('CreateCampaignMemberComponent', () => {
  let component: CreateCampaignMemberComponent;
  let fixture: ComponentFixture<CreateCampaignMemberComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateCampaignMemberComponent
      ],
      providers: [
        CampaignMemberService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateCampaignMemberComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});