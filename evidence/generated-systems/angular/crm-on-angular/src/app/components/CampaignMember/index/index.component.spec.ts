
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexCampaignMemberComponent } from './index.component';
import { CampaignMemberService } from '../../../services/CampaignMember.service';

describe('IndexCampaignMemberComponent', () => {
  let component: IndexCampaignMemberComponent;
  let fixture: ComponentFixture<IndexCampaignMemberComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexCampaignMemberComponent
      ],
      providers: [
        CampaignMemberService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexCampaignMemberComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});