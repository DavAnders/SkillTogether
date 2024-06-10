import PropTypes from "prop-types";
import { useState } from "react";
import UpdateInterest from "./UpdateInterest";
import "../styles/SkillItem.css";

function InterestItem({ interest, handleDelete, onInterestUpdated }) {
  const [isEditing, setIsEditing] = useState(false);

  return (
    <div className="interest-item">
      <li key={interest.id}>
        {interest.interest}
        {!isEditing && (
          <div
            style={{ display: "flex", gap: "10px" }}
            className="interest-buttons"
          >
            <button onClick={() => handleDelete(interest.id)}>Delete</button>
            <button onClick={() => setIsEditing(true)}>Edit</button>
          </div>
        )}
        {isEditing && (
          <UpdateInterest
            interestId={interest.id}
            onInterestUpdated={() => {
              setIsEditing(false);
              onInterestUpdated();
            }}
          />
        )}
      </li>
    </div>
  );
}

InterestItem.propTypes = {
  interest: PropTypes.shape({
    id: PropTypes.number,
    interest: PropTypes.string,
  }).isRequired,
  handleDelete: PropTypes.func.isRequired,
  onInterestUpdated: PropTypes.func.isRequired,
};

export default InterestItem;
